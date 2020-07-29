

class IncomingEventDetail(object):

    def __init__(self, d):
        self.key = d['key']


class IncomingEvent(object):

    def __init__(self, d):
        self.id = d['id']
        self.type = d['detail-type']

        self.detail = IncomingEventDetail(d['detail'])
