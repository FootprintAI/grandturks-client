#### python cli ####

This python wrapper is wrapping the underlying cli binary to support functionalities:

- create a project
- upload a file to the project's bucket
- create a datasource with the uploaded file
- create model inference with a system kubeflow pipeline 
- create a job to binding the datasource and the model inferencer

so you can just to ahead to dashboard to check the inference results.

#### Usage

```
ACCOUNT=<your-account> PASSWORD=<your-password> MODEL_URI=<model-uri> python3 cli.py --alsologtostderr
```