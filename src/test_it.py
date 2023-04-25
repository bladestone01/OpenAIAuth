
from OpenAIAuth import Authenticator

if __name__ == '__main__':
    auth = Authenticator("65040491@qq.com", "_xxxxx", "tencent@ai123")

    auth.begin()
    access_token = auth.get_access_token()
    print("hello world=>" + access_token)