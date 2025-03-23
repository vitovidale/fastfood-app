fastfood-app
Aplicação de Autoatendimento FastFood

📄 Descrição do Projeto
O fastfood-app é a interface de autoatendimento que permite aos clientes montar pedidos (lanche, acompanhamento, bebida, sobremesa), visualizar o status e confirmar pagamentos. Faz parte do sistema FastFood‑App integrado com autenticação (Auth0), banco de dados (RDS) e backend em Kubernetes (EKS).

⚙️ Tech Stack
Linguagem: Go

Container: Docker

Orquestração: Kubernetes (EKS)

Registro de imagem: AWS ECR

Banco de Dados: Amazon RDS (PostgreSQL)

Autenticação: API Gateway + Lambda (Auth0)

CI/CD: GitHub Actions

⚠️ Pré‑requisitos
Docker

kubectl configurado para o cluster EKS

AWS CLI com credenciais (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)

Variáveis de ambiente definidas em .env (veja .env.sample)

🔧 Configuração Local
Clone o repositório:

bash
Copiar
git clone https://github.com/vitovidale/fastfood-app.git
cd fastfood-app
Copie o exemplo de variáveis:

bash
Copiar
cp .env.sample .env
Build & run via Docker:

bash
Copiar
docker build -t fastfood-app .
docker run --env-file .env -p 3000:3000 fastfood-app
Acesse em http://localhost:3000

📡 Endpoints Principais
Método	Rota	Descrição
GET	/menu	Lista todos os produtos
POST	/orders	Cria novo pedido
GET	/orders/{id}	Consulta status de um pedido
✅ Testes
Se houver testes automatizados:

bash
Copiar
go test ./...
🚀 Deploy
O deploy da aplicação é feito automaticamente via GitHub Actions (workflow_dispatch).

Como executar
No GitHub, abra este repositório → clique na aba Actions

Selecione o workflow Deploy FastFood App

Clique em Run workflow

📦 Kubernetes
Deployment: fastfood-app

Namespace: default

Para rollback:

bash
Copiar
kubectl rollout undo deployment/fastfood-app -n default
🛠 Troubleshooting
Erro	Solução rápida
ImagePullBackOff	Verifique se a imagem foi enviada ao ECR corretamente
DB connection refused	Confirme string de conexão e regras do Security Group
Unauthorized (401)	Verifique variáveis Auth0 (CLIENT_ID/SECRET)
🤝 Contribuição
Fork → feature branch → PR para develop

Branch protegida → revisão obrigatória

Convenção de commits: Conventional Commits

📜 Licença
MIT License — veja o arquivo LICENSE para detalhes.
