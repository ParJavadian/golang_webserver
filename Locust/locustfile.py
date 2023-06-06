from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        self.client.get("/get_users_sql?user_id=5&message_id=0&auth_key=13",verify=False)
        # self.client.get("/world")