"""eh"""
import requests

with open("./anilist_page.gql", "r", encoding="utf-8") as f:
    gql_query = f.read()

CSV_FILE = "romanji,img_url,anilist_id\n"
SQL_FILE = "INSERT INTO daily (romanji, img_url, anilist_id) VALUES\n"
for page in range(1, 7):
    response = requests.post(
        "https://graphql.anilist.co",
        json={
            "query": gql_query,
            "variables": {"page": page, "perPage": 50},
        },
        timeout=5000,
    )

    for m in response.json()["data"]["Page"]["media"]:
        title = m["title"]["english"] or m["title"]["romaji"]
        title = title.replace("'", "''")
        CSV_FILE += f'{title}, {m["coverImage"]["medium"]}, {m["id"]}\n'
        SQL_FILE += (
            f'\t(\'{title}\', \'{m["coverImage"]["medium"]}\', \'{m["id"]}\'),\n'
        )

with open("./sql_dailies.txt", "w", encoding="utf-8") as sql_file:
    sql_file.write(SQL_FILE)

# print(CSV_FILE)
print("DONE")
