import uuid


def get_uuid():
    uid = str(uuid.uuid4())
    return ''.join(uid.split('-'))


if __name__ == '__main__':
    print(get_uuid())
