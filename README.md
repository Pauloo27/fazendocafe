# Fazendo Café

Servidor de páginas `...ndo.cafe`, usando Go, um certificado SSL com wildcard 
e um template HTML para serve diversas páginas como subdomínios de `.ndo.cafe`.

## Páginas Disponíveis

Você pode ver a [lista completa em ndo.cafe](https://ndo.cafe).

- [quere.ndo.cafe](https://quere.ndo.cafe)
- [faze.ndo.cafe](https://faze.ndo.cafe)
- [toma.ndo.cafe](https://toma.ndo.cafe)
- [danca.ndo.cafe](https://danca.ndo.cafe)
- [bebe.ndo.cafe](https://bebe.ndo.cafe)
- [come.ndo.cafe](https://come.ndo.cafe)
- [derrama.ndo.cafe](https://derrama.ndo.cafe)
- [queima.ndo.cafe](https://queima.ndo.cafe)

## Sugerir Página

Se você tem alguma sugestão que termina com `.ndo.cafe`, você pode enviar um 
Pull Request implementado a página ou, se preferir, criar uma Issue e esperar
alguém implementar para você.

Antes de fazer qualquer uma das 2, tenha:

1 - o texto: será usado como URL
1 - uma imagem (GIF, PNG ou JPEG)
1 - um título
1 - um subtítulo
1 - uma descrição (será usada nas metadatas da página)

_uma sugestão é apenas uma sugestão e precisa ser aceita para que seja 
adicionada_

### Contribuindo

Adicione a imagem na pasta `web/asset/` e então adicione a página no arquivo
`internal/server/pages.go`.

Para ver como a página fica, inicie o servidor na sua máquina:

Copie o arquivo `example.env` para `.env` e inicie o servidor com `make run`. 
Abra no navegador `http://localhost:1235/sugeri.ndo.cafe`. 
(onde `sugeri.ndo.cafe` é a página que você quer ver).

## Licença

<img src="https://i.imgur.com/AuQQfiB.png" alt="GPL Logo" height="100px" />

Este projeto é licenciado sob [GNU General Public License v2.0](./LICENSE).
