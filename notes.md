recebe mensagem via rabbit para saber qual vidio converter
faz o download do video no google cloud storage
fragmenta o video 
converte para MPEG-DASH
faz o upload do video para o google cloud storage
envia uma notificacao via filas com as informacoes do vide para converter ou infarmando erro na conversao
em caso de erro a mensagem original enviada via rabbit sera rejeitafa e encaminhada a uma dead letter exchange


Arquitetura do projeto HEXAGONAL

´´´ yaml
- Framework
  - Application
    - Domain
´´´

formato mensagens

input
``` json
{
    "resource_id": "qualquer coisa, qualquer formato",
    "file_path": "convite.mp4"

}
```

output
```json
{
    "id": "algum id",
    "output_bucket_path": "/alguma_coisa.mp4",
    "status": "sucesso",
    "video" : {
        "encoded_video_folder": "/folder",
        "reource_id": "id original",
        "file_path": "path original",
    },
    "error": "",
    "created_path" : "timestamp",
    "updated_at": "timestamp"
}
```


error
```json
{
    "message": {
        "resource_id": "qualquer coisa, qualquer formato",
        "file_path": "convite.mp4"
    },
    "error": "reason"
}
```
