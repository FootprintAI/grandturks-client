"""python version cli for running grandturks"""

import logging

LOCAL_BIN_PATH = "./cli"
DEMO_PROJECT_NAME = "sp8-demo-project"
STORAGE_PREFIX = "demo/realtimequery"
STORAGE_QUERY_OBJECT_NAME = "forquery"
STORAGE_QUERY_OBJECT_INDEX= "objectindex"
DEFAULT_MODEL_NAME = "carreid13032"
DEFAULT_CLI_VERSION="v1.2.3"

def load_json(output: str) -> list:
    import json
    if len(output) == 0:
        return []
    return json.loads(output)

def ensure_cli():
    import os
    import platform
    import wget

    if os.path.exists(LOCAL_BIN_PATH):
        return

    logging.info('local cli does not exist, downloading it from web')
    sys_name = platform.system().lower()
    cli_uri = """https://github.com/FootprintAI/grandturks-client/releases/download/{0}/grandturks-cli.{1}""".format(DEFAULT_CLI_VERSION, sys_name)

    wget.download(cli_uri, LOCAL_BIN_PATH)
    ensure_runable(LOCAL_BIN_PATH)
    logging.info('local cli downloaded')

def ensure_runable(local_bin_path: str):
    import os
    import stat

    st = os.stat(local_bin_path)
    os.chmod(local_bin_path, st.st_mode | stat.S_IEXEC)

def _exec(args: str) -> str:
    import subprocess

    args_list = args.split()
    args_list.insert(0, LOCAL_BIN_PATH)

    logging.info('run with commands: %s', " ".join(args_list))

    popen = subprocess.Popen(args_list, stdout=subprocess.PIPE)
    popen.wait()
    output = popen.stdout.read()
    return output

def login(email_account: str, account_password: str):
    _exec("""
    login basic --email={0} --password={1}
     """.format(email_account, account_password))


def ensure_demo_project_exists(create_project_if_not_found = True) -> dict :
    """ ensure_demo_project_exists ensures demo project exists and 
    return its project id
    """
    list_project_json = load_json(_exec("""
    list project --format=json
    """))
    for project in list_project_json:
        if project["name"] == "sp8-demo-project":
            return project
    # project not found, create one
    _exec("""
    create project --name=sp8-demo-project --desc=for demo --format=json
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
    upload storage --project_id={0} --bucket_name={1} --prefix={2} {3}
    """.format(project_id, project_private_bucket, prefix, src_path))
 

def ensure_datasource_exists(project: dict) -> dict :
    """ ensure_datasource_exists ensures data source exists and returns its
    data source id"""
    import os

    target_index_object = os.path.join(STORAGE_PREFIX, STORAGE_QUERY_OBJECT_INDEX)
    project_id = project['id']
    project_private_bucket = project['private_bucket']
    list_datasource_json = load_json(_exec("""
    list datasource --project_id={} --format=json
    """.format(project_id)))
    for datasource in list_datasource_json:
        if datasource["index_object"] == target_index_object:
            return datasource

    # datasource not found, create one
    _exec("""
    create datasource --project_id={0} --bucket={1} --format=json
    --index_object={2}
    """.format(project_id, project_private_bucket, target_index_object))
    return ensure_datasource_exists(project)

def ensure_inference_exists(project: dict, model_uri: str):
    project_id = project['id']
    list_inference_json = load_json(_exec("""
    list inference --project_id={} --format=json
    """.format(project_id)))
    for inference in list_inference_json:
        if inference["model_name"] == DEFAULT_MODEL_NAME:
            return inference

    # datasource not found, create one
    named_pipeline = get_named_pipeline(project)
    _exec("""
    create inference --project_id={0} --named_pipeline_id={1}
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
            get inference --project_id={0} --inference_id={1} --format=json
            """.format(project_id, inference_id)))
        if runtime_inference[0]["model_status"] == "true":
            break
        logging.info('model is not ready, try for next 5s')
        time.sleep(5)

def get_named_pipeline(project:dict) -> dict:
    """ get_named_pipeline get the target name pipeline and return its id"""
    project_id = project['id']
    list_pipeline_json = load_json(_exec("""
    list pipeline --project_id={0} --format=json
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
    create job --project_id={0} --inference_id={1} --datasource_id={2}
    --output_tsdb_name=p13carreid --output_stream_name=p13carreid
    --outout_storage_prefix=p13carreid --format=json
    """.format(project_id, inference_id, datasource_id))

def run_demo(target_filename: str):
    import os

    model_uri = os.getenv('MODEL_URI')
    email_account = os.getenv('ACCOUNT')
    accont_password = os.getenv('PASSWORD')

    ensure_cli()
    login(email_account, accont_password)
    project = ensure_demo_project_exists()
    src_dir = prepare_uploading_dir(target_filename)
    upload_file_to_bucket(project, STORAGE_PREFIX, src_dir)
    datasource = ensure_datasource_exists(project)
    inference = ensure_inference_exists(project, model_uri)
    # we have to wait unitl the model is ready
    wait_until_inference_is_ready(project, inference)
    create_job_running(project, datasource, inference)
    logging.info('you are goood to open grafana page')

if __name__ == '__main__':
    import argparse
    import sys

    parser = argparse.ArgumentParser()
    parser.add_argument('--alsologtostderr', dest='alsologtostderr',
            default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--logfile', '-l', dest='logfile', type=str, default='cli-py.log')
    parser.add_argument('--target_filename', '-t', dest='target_filename', type=str, default='2.jpg')
    args = parser.parse_args()


    logformat = '%(asctime)s - %(name)s - %(levelname)s - %(message)s'
    logging.basicConfig(filename=args.logfile, level=logging.INFO, format=logformat)
    if args.alsologtostderr:
        rootlogger = logging.getLogger()
        handler = logging.StreamHandler(sys.stdout)
        handler.setLevel(logging.INFO)
        handler.setFormatter(logging.Formatter(logformat))
        rootlogger.addHandler(handler)
    run_demo(args.target_filename)

