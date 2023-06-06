from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        self.client.get("/req_pq?nonce=hihihi&message_id=0",verify=False)
        # self.client.get("/world")