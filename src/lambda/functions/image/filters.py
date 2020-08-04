from timeit import default_timer as timer

from PIL.Image import core as _imaging
from io import BytesIO


def monochrome(buf):
    start = timer()

    img = _imaging.open(BytesIO(buf))
    img = img.convert("L")

    buffer = BytesIO()
    img.save(buffer, 'PNG')
    buffer.seek(0)

    print("Image converted: latency={}".format(timer()-start))

    return buffer
