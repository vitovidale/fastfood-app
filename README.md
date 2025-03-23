# fastfood-app

> Aplicação de Autoatendimento FastFood

---

## Descrição do Projeto

O **fastfood-app** é a interface de autoatendimento que permite aos clientes montar pedidos (lanche, acompanhamento, bebida, sobremesa), visualizar o status e confirmar pagamentos. Faz parte do sistema FastFood‑App integrado com autenticação (Auth0), banco de dados (RDS) e backend em Kubernetes (EKS).

---

## Pré‑requisitos

- Docker  
- kubectl configurado para o cluster EKS  
- AWS CLI com credenciais (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)  
- Variáveis de ambiente definidas em `.env` (veja `.env.sample`)

---

## 🚀 Deploy

O deploy é feito automaticamente via GitHub Actions (workflow_dispatch).

### Como executar

1. No GitHub, abra este repositório → clique em **Actions**  
2. Selecione o workflow **Deploy FastFood App**  
3. Clique em **Run workflow**

---
