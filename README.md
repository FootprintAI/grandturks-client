# Kafeido

Kafeido is an AI platform, which is built on top of kubeflow and kubernetes, to provide real-time data streaming inferences with multiple machine learning models in low-energy consumption, low hardware spec requirement, and trustworthy ability.

Kafeido, a nicked name for kaleido, is a tunnel which wires datasources with machine learning models and support the following mapping:

* One-datasource-one-model
* One-datasource-multiple-models
* Multiple-datasources-one-model
* multiple-datasources-multiple-models

And we believe that this could serve the most applications right now we are facing.

Kafeido can be deployed on top of [multikf](https://github.com/footprintai/multikf) (a command line library to provision kubernetes cluster on the single host machine) or be deployed with [kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/) which is kubernetes’s official deployment command line tool for kubernetes. Please [contact us](https://get-tintin.footprint-ai.com) if you want to provision Kafeido on your own machines.

The minimum hardware requirements for server is 8 Core+ CPU, 18G+ system memory, 50G+ Disk spaces, and GPU accelerators. There are no hardware requirements for its client.


## Usage

#### Create Project

```
create project --name $projectname \
--desc "project for demonstration"
```

#### Create Datasource

```
./cli create datasource --project_id=$project_id \
 --bucket_name=$project_bucket \
 --index_object="videos/<some-video-prefix>.mp4" \
 --duration_per_frame=1s \
 --fps=1 \
 --type=VIDEO
```

#### Create Model

```
./cli create inference --project_id=$project_id \
 --bucket_name=$project_bucket \
 --model_uri=kafeido:///$project_bucket/$object_path
```

#### Create Prediction

```
 ./cli create predict --project_id=$project_id \
 --inference_id=$inference_id \
 --query_file=example.wav --query_type=audio --query_lang=en
```