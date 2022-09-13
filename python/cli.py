"""python version cli for running grandturks"""

DEMO_PROJECT_NAME = "sp8-demo-project"
STORAGE_PREFIX = "demo/realtimequery"
STORAGE_QUERY_OBJECT_NAME = "forquery"
STORAGE_QUERY_OBJECT_INDEX= "objectindex"
DEFAULT_MODEL_NAME = "carreid13032"

def load_json(output: str) -> list:
    if len(output) == 0:
        return []
    import json
    return json.loads(output)

def _exec(args: str) -> str:
    print(args)
    import subprocess
    popen = subprocess.Popen(args.split(), stdout=subprocess.PIPE)
    popen.wait()
    output = popen.stdout.read()
    return output

def login(email_account: str, account_password: str):
    _exec("""
    ./cli login basic --email={0} --password={1}
     """.format(email_account, account_password))


def ensure_demo_project_exists(create_project_if_not_found = True) -> dict :
    """ ensure_demo_project_exists ensures demo project exists and 
    return its project id
    """
    list_project_json = load_json(_exec("""
    ./cli list project --format=json
    """))
    for project in list_project_json:
        if project["name"] == "sp8-demo-project":
            return project
    # project not found, create one
    _exec("""
    ./cli create project --name=sp8-demo-project --desc=for demo --format=json
    """)
    return ensure_demo_project_exists(False)



def prepare_uploading_dir(target_filename_for_query: str) -> str:
    """ prepare_uploading_dir create a temp dir and prepare two files inside
    1. objectindex file contains a single line of file name which is
    demo/realtimequery/forquery
    2. demo/realtimequery/forquery is the $target_filename_for_query itself
    with cp cli

    and return the temp dir
    """
    import tempfile
    import shutil
    import os
    tempdir = tempfile.mkdtemp()
    # create objectindex file
    with open(os.path.join(tempdir, "objectindex"), 'w+') as fd:
        fd.write("{}/{}\n".format(STORAGE_PREFIX, STORAGE_QUERY_OBJECT_NAME))
    shutil.copy2(target_filename_for_query, os.path.join(tempdir, STORAGE_QUERY_OBJECT_NAME))
    return tempdir

def upload_file_to_bucket(project: dict, prefix: str, src_path: str):
    project_id = project['id']
    project_private_bucket = project['private_bucket']
    _exec("""
    ./cli upload storage --project_id={0} --bucket_name={1} --prefix={2} {3}
    """.format(project_id, project_private_bucket, prefix, src_path))
 

def ensure_datasource_exists(project: dict) -> dict :
    """ ensure_datasource_exists ensures data source exists and returns its
    data source id"""
    import os

    target_index_object = os.path.join(STORAGE_PREFIX, STORAGE_QUERY_OBJECT_INDEX)
    project_id = project['id']
    project_private_bucket = project['private_bucket']
    list_datasource_json = load_json(_exec("""
    ./cli list datasource --project_id={} --format=json
    """.format(project_id)))
    for datasource in list_datasource_json:
        if datasource["index_object"] == target_index_object:
            return datasource

    # datasource not found, create one
    _exec("""
    ./cli create datasource --project_id={0} --bucket={1} --format=json
    --index_object={2}
    """.format(project_id, project_private_bucket, target_index_object))
    return ensure_datasource_exists(project)

def ensure_inference_exists(project: dict, model_uri: str):
    project_id = project['id']
    list_inference_json = load_json(_exec("""
    ./cli list inference --project_id={} --format=json
    """.format(project_id)))
    for inference in list_inference_json:
        if inference["model_name"] == DEFAULT_MODEL_NAME:
            return inference

    # datasource not found, create one
    named_pipeline = get_named_pipeline(project)
    _exec("""
    ./cli create inference --project_id={0} --named_pipeline_id={1}
    --model_name={2} --model_uri={3} --format=json
    """.format(project_id, named_pipeline["named_pipeline_id"],
        DEFAULT_MODEL_NAME, model_uri))
    return ensure_inference_exists(project, model_uri)

def wait_until_inference_is_ready(project: dict, inference: dict, max_retries = 100):
    import time

    project_id = project['id']
    inference_id = inference['model_inference_id']
    for try_count in range(0, max_retries):
        runtime_inference = load_json(_exec("""
            ./cli get inference --project_id={0} --inference_id={1} --format=json
            """.format(project_id, inference_id)))
        if runtime_inference[0]["model_status"] == "true":
            break
        print("model is not ready, try for next 5s")
        time.sleep(5)

def get_named_pipeline(project:dict) -> dict:
    """ get_named_pipeline get the target name pipeline and return its id"""
    project_id = project['id']
    list_pipeline_json = load_json(_exec("""
    ./cli list pipeline --project_id={0} --format=json
    """.format(project_id)))
    for pipeline in list_pipeline_json:
        if pipeline["named_pipeline_name"] == "kserve:car-reid":
            return pipeline
    raise Exception("no such pipeline")

def create_job_running(project: dict, datasource: dict, inference:dict):
    project_id = project['id']
    datasource_id = datasource['datasource_id']
    inference_id = inference['model_inference_id']
    _exec("""
    ./cli create job --project_id={0} --inference_id={1} --datasource_id={2}
    --output_tsdb_name=p13carreid --output_stream_name=p13carreid
    --outout_storage_prefix=p13carreid --format=json
    """.format(project_id, inference_id, datasource_id))

def run_demo(target_filename: str):
    import os
    model_uri = os.getenv('MODEL_URI')
    email_account = os.getenv('ACCOUNT')
    accont_password = os.getenv('PASSWORD')

    login(email_account, accont_password)
    project = ensure_demo_project_exists()
    src_dir = prepare_uploading_dir(target_filename)
    upload_file_to_bucket(project, STORAGE_PREFIX, src_dir)
    datasource = ensure_datasource_exists(project)
    inference = ensure_inference_exists(project, model_uri)
    # we have to wait unitl the model is ready
    wait_until_inference_is_ready(project, inference)
    create_job_running(project, datasource, inference)
    print("you are goood to open grafana page")

if __name__ == '__main__':
    target_filename = "2.jpg"

    run_demo(target_filename)

