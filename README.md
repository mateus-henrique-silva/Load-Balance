# Load Balance API

## Dev

Para executar a api é necessario utilizar do comando abaixo na pasta raiz do projeto.

```air```

## Como Funciona ?

A api se trata de um codigo que utiliza de um middleware que verifica o tamanho das requisicoe e caso uma esteja com o processamento sobrecarregado ele redireciona para o outro servidor.

Imagine que você é dono de uma sorveteria e tem várias pessoas esperando para comprar sorvetes. Você é o único atendente, e cada cliente tem um pedido diferente. Alguns querem picolés, outros milkshakes, e alguns preferem sorvetes em casquinhas.

Agora, pense que cada cliente é como uma solicitação que chega ao seu servidor de internet. Cada tipo de sorvete que você vende é como um serviço diferente que seu servidor precisa fornecer, como uma página da web, um vídeo ou uma imagem.

Se você é o único atendente, pode ficar difícil lidar com todos os pedidos ao mesmo tempo, especialmente se alguns forem mais complicados ou demorados de preparar. Aqui é onde entra o "load balancing" (equilíbrio de carga).

O "load balancing" é como ter mais atendentes trabalhando na sorveteria. Eles ajudam a distribuir os clientes de forma equitativa, para que todos sejam atendidos de maneira eficiente e ninguém fique esperando muito tempo.

No mundo da internet, o "load balancing" faz algo semelhante. Ele distribui as solicitações dos usuários entre vários servidores, garantindo que nenhum servidor fique sobrecarregado. Isso melhora o desempenho, a confiabilidade e a velocidade dos serviços online, garantindo que todos os usuários recebam o que precisam de maneira rápida e eficiente.

Em resumo, o "load balancing" é como ter atendentes extras na sua sorveteria da internet, garantindo que todos os clientes (usuários) sejam atendidos de maneira rápida e eficiente, sem sobrecarregar nenhum servidor.