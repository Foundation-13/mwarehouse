from PIL import Image
from io import BytesIO
from timeit import default_timer as timer


def monochrome(buf):
    start = timer()

    img = Image.open(BytesIO(buf))
    img = img.convert("L")

    buffer = BytesIO()
    img.save(buffer, 'PNG')
    buffer.seek(0)

    print("Image converted: latency={}".format(timer()-start))

    return buffer
