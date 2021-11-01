## Saiga Demo

## Env File
Please change AWS_SECRET_KEY and AWS_SECRET_ACCESS_KEY in .env file or copy your .~/aws/credentials to docker

## How to run it ? 
 - `cd build`
 - `docker-compose up`

Note: if you just want to run mongodb through docker and want to run the API through the localhost
 - `change` "mongodb://mongodb:27017" `to` "mongodb://localhost:27017"

## Endpoints
`/api/user/login` : You can login with this endpoint example: ` {"email":"customer@gmail.com", "password":"123456"}`
`/api/user/signup` :  You can register with this endpoint example: ` {"email":"customer@gmail.com", "password":"123456"}`
`/api/welcome` :  Just welcome page.
`/api/tasks`: You can get all tasks if you are a operator
`/api/chats/{id}/messages` : You can get all message of the chat if you are a operator
`/api/attachments/{id}/download` :  You can download attachments belonging to chat if you are a operator

Note: if you added role on the customer you can use role field example: ` {"email":"customer@gmail.com", "password":"123456", "role":"operator"}` default value is "customer"


## How to deploy to ECS?

    - Login AWS and use Amazon ECR service create private repo
    - docker tag build_app <account-id>.dkr.ecr.<region>.amazonaws.com/build_app
    - docker push <account-id>.dkr.ecr.<region>.amazonaws.com/build_app
    - Configuring ECS cluster
        - EC2 Linux + Networking
        - Cluster Name = Saiga
        - On-Demand Instance
        - t2.small
        - Number of instance = 1
        - EBS Store = 10
        - Create Networking
        - Create IAM
        - Create Task Definitions
            - Network = Bridge
            - ecsTaskExecutionRole
            - TaskMemory = 250MB
            - TaskCPU = 128
    - Creating ALB
        - Listener HTTP: 80
    - Configure Routing
        - Port 8080
 
## How to build dockerfile ?
- docker-compose build
 
![alt text](https://uploads-ssl.webflow.com/6110ab2a9f23df3a7351476a/6127bb9e0ffb8dcd9171d33b_LOGO.svg)

  