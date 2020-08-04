from timeit import default_timer as timer

from PIL.Image import open
from io import BytesIO


def monochrome(buf):
    start = timer()

    img = open(BytesIO(buf))
    img = img.convert("L")

    buffer = BytesIO()
    img.save(buffer, 'PNG')
    buffer.seek(0)

    print("Image converted: latency={}".format(timer()-start))

    return buffer
