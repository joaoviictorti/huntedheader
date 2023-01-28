<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=header"/>

[![Typing SVG](https://readme-typing-svg.herokuapp.com/?color=0000FF&size=32&center=true&vCenter=true&width=1000&height=30&lines=huntedheader)](https://git.io/typing-svg)



<h4 align="center"> Tool to generate payloads focused on CSRF </h4>


<p align="center">
  <a href="#características">Features</a> •
  <a href="#instalação">Install</a> •
  <a href="#forma-de-utilização">How to use</a> •
  <a href="#executando-revshell">Usage</a>
</p>

---


O huntedheader é uma ferramenta que tem o objetivo de verificar se é existente os cabeçalhos de segurança numa resposta HTTP/HTTPS de uma url mencionada.

Projetei o `huntedheader` e mantive um modelo consistentemente passivo para torná-lo útil para testadores de penetração.

# Características

 - Verificar se é existente os cabeçalhos de segurança numa resposta HTTP/HTTPS

# Forma de utilização

```sh
huntedheader -u https://teste.com
```
Isso exibirá a ajuda para a ferramenta. Aqui estão todos os switches que ele suporta:
```yaml
usage: Hunteheader [-h|--help] [-u|--url "<value>"] [-f|--file "<value>"]      


    __                __           ____                   __
   / /_  __  ______  / /____  ____/ / /_  ___  ____ _____/ /__  _____
  / __ \/ / / / __ \/ __/ _ \/ __  / __ \/ _ \/ __ / __  / _ \/ ___/
 / / / / /_/ / / / / /_/  __/ /_/ / / / /  __/ /_/ / /_/ /  __/ /
/_/ /_/\__,_/_/ /_/\__/\___/\__,_/_/ /_/\___/\__,_/\__,_/\___/_/



Arguments:

  -h  --help  Print help information
  -u  --url   Insert url. Default: url
  -f  --file  Insert file url. Default: file

```

# Instalação

huntedheader requer o **golang** instalado, para usar:

```sh
go install -v github.com/huntedheader@latest
```

# Executando huntedheader

```console
huntedheader -u https://teste.com
[28/01/2023 00:11:30] [X-Content-Type-Options] https://teste.com
[28/01/2023 00:11:30] [Referrer-Policy] https://teste.com
[28/01/2023 00:11:30] [Permissions-Policy] https://teste.com     
[28/01/2023 00:11:30] [Strict-Transport-Security] https://teste.com

```


<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=footer"/>
