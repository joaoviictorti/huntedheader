<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=header"/>

![hunteheader](/img/hunterheader.png)

<p align="center">
	<a href="https://www.python.org/"><img src="https://img.shields.io/badge/made%20with-go-blue"></a>
	<a href="#"><img src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-blueviolet"></a>
	<a href="https://github.com/joaoviictorti/hunterheader/releases"><img src="https://img.shields.io/github/release/joaoviictorti/hunterheader?color=blue"></a>
</p>

<h4 align="center">Hunteheader verificar os cabeçalhos de segurança ausente em uma resposta HTTP / HTTPS</h4>


<p align="center">
  <a href="#características">Características</a> •
  <a href="#instalação">Instalação</a> •
  <a href="#forma-de-utilização"> Forma de utilização</a> •
  <a href="#detalhes">Detalhes</a> •
  <a href="#executando-revmap">Executando revmap</a>  
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

# Detalhes

![hunteheader](/img/help.png)


# Instalação

huntedheader requer o **golang** instalado, para usar:

```sh
go install -v github.com/joaoviictorti/huntedheader/cmd/hunterheader@latest
```

# Executando huntedheader

![hunteheader](/img/exec.png)


<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=footer"/>
