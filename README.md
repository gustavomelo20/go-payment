# API de Processamento de Pagamentos

![Processamento de Pagamentos](https://i.imgur.com/cEowwPz.png)

## Visão Geral

Esta API processa pagamentos via diferentes formas, como Pix e Cartão de Crédito, de forma concorrente, retornando resultados detalhados para cada pagamento.

---

## Endpoint da API

**POST** `http://localhost:8080/charge`

---

## Exemplo de Corpo da Requisição

```json
{
  "charge_id": "5eeb3f82-b380-47a3-97c9-d02de241daa7",
  "payments": [
    {
      "method": "PIX",
      "amount": 150.00,
      "details": {
        "pix_key": "teste@exemplo.com"
      }
    },
    {
      "method": "CREDIT_CARD",
      "amount": 200.00,
      "details": {
        "card_number": "4111111111111111",
        "expiry": "12/26",
        "cvv": "123"
      }
    }
  ]
}
```

## Exemplo de Corpo da Requisição

```json
{
  "charge_id": "5eeb3f82-b380-47a3-97c9-d02de241daa7",
  "results": [
    {
      "payment_id": "aef0e4a3-3cdb-4218-8f6b-9c44d84fa877",
      "status": "APPROVED",
      "message": "Pagamento com cartão de crédito aprovado com sucesso",
      "processed_at": "2025-06-09T15:01:29.468237-03:00",
      "details": {
        "authorization_number": "CC-1749492089",
        "card_last4": "1111"
      }
    },
    {
      "payment_id": "e06725fd-c0ef-4304-96f7-8d2ef592ce52",
      "status": "APPROVED",
      "message": "Pagamento Pix aprovado com sucesso",
      "processed_at": "2025-06-09T15:01:29.468237-03:00",
      "details": {
        "pix_key": "teste@exemplo.com",
        "receipt_number": "PIX-1749492089"
      }
    }
  ]
}
```