JWT_DIR=config/jwt
PRIVATE_KEY=$(JWT_DIR)/private.pem
PUBLIC_KEY=$(JWT_DIR)/public.pem

.PHONY: generate-jwt-keys

generate-jwt-keys:
	@echo "🔐 Gerando chaves JWT (se não existirem)"
	@[ -d $(JWT_DIR) ] || mkdir -p $(JWT_DIR)
	@[ -f $(PRIVATE_KEY) ] || openssl genpkey -algorithm RSA -out $(PRIVATE_KEY) -pkeyopt rsa_keygen_bits:2048
	@[ -f $(PUBLIC_KEY) ]  || openssl rsa -pubout -in $(PRIVATE_KEY) -out $(PUBLIC_KEY)