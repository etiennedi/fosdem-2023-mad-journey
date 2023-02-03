import weaviate
from typing import Optional
import uuid
from loguru import logger

reviews = [
        {
            "varietal": "Riesling",
            "vintage": 2017,
            "name": "Koehler Ruprecht Saumagen Spaetlese R trocken ",
            "review": "As expected, this 2017 is brutally young. Immediately after opening it gives off hints as how it will develop. As the KR team have confirmed to me, this could be quite close to the 2009 which is one of my all time favorite. Nevertheless, this needs a lot of time. It shuts down with air, rather than opening up. I hope I have to patience to wait at least 5 years to open the next bottle. I'm convinced it will show the this perfect balance of creamy acididty and flavor complexity that Koehler Ruprecht is known for. It certainly has plenty of acidity to age for a very long time. Luckily, I have quite a few.",
        },
        {
            "varietal": "Riesling",
            "vintage": 2009,
            "name": "Koehler Ruprecht Saumagen Auslsese R trocken",
            "review": "The 2009 Auslese R trocken is probably one of my favorite wines from KR – if not favorite of all Rieslings out there. This vintage really shows the difference between the ever so slightly leaner Spaetlese and the bigger and bolder Auslese. Yet, now in its perfect drinking window, it keeps an amazing freshness through its acidity. Wild, yet creamy with beautiful notes of Mango yoghurt, pear, and apricot. But of course it keeps the telltale minerality that makes Saumagen Saumagen.",
        },
        {
            "varietal": "Nebbiolo",
            "vintage": 2017,
            "name": "Dosio Barolo Bussia",
            "review": "I was surprised that a 2017 would be drinkable already. This Barolo is a great example of why the Nebbiolo grape is so highly regarded. It is definitely not made in a style that needs decades before the tart tannins become smooother. It is pretty smooth right now already. We used this at a team event in the Piemont region and it was a perfect way to introduce everyone to what makes this region so special – without breaking the bank.",
        },
]

def reset_schema(client: weaviate.Client):
    client.schema.delete_all()
    class_obj = {
        "vectorizer": "text2vec-openai",
        "class": "Review",
        "properties": [
            {
                "dataType": [ "string" ],
                "name": "varietal",
            },
            {
                "dataType": [ "string" ],
                "name": "name",
            },
            {
                "dataType": [ "int" ],
                "name": "vintage",
            },
            {
                "dataType": [ "text" ],
                "name": "review",
            },
        ]
    }

    client.schema.create_class(class_obj)

def load_records(client: weaviate.Client, objs):
    client.batch.configure(batch_size=100, callback=handle_errors)
    i = 0
    with client.batch as batch:
        for obj in objs:
            batch.add_data_object(
                data_object=obj,
                class_name="Review",
                uuid=uuid.UUID(int=i),
            )
            i+=1
    logger.info(f"Finished writing {len(objs)} records")


def handle_errors(results: Optional[dict]) -> None:
    """
    Handle error message from batch requests logs the message as an info message.
    Parameters
    ----------
    results : Optional[dict]
        The returned results for Batch creation.
    """

    if results is not None:
        for result in results:
            if (
                'result' in result
                and 'errors' in result['result']
                and 'error' in result['result']['errors']
            ):
                for message in result['result']['errors']['error']:
                    logger.error(message['message'])



client = weaviate.Client("http://localhost:8080")
reset_schema(client)
load_records(client, reviews)
