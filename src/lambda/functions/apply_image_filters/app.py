from model import events


def lambda_handler(event, context):
    """
    """

    parsed_event = events.IncomingEvent(event)
    print("Received event({}, {}, {})".format(parsed_event.id, parsed_event.type, parsed_event.detail.key))

    return parsed_event.detail.key

