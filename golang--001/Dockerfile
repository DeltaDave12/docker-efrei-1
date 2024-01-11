# LE BUT DE CE DOCKERFILE : Compiler le language et installer dépendance
# Étape de construction
#FROM golang:latest as builder
# Alpine est un OS étant désigner pour être le plus petit possible
# FORM golang:alpine

#WORKDIR /go/src/app

# Copiez uniquement les fichiers nécessaires pour la construction
#COPY go.mod go.sum ./
#RUN go mod download

#COPY . .

# Construisez votre application
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Étape de production
#FROM scratch

# Copiez l'exécutable construit depuis l'étape de construction
#COPY --from=builder /go/src/app/main /app/main

# Définissez le répertoire de travail pour l'exécution
#WORKDIR /app

# Définissez la commande à exécuter lorsque le conteneur démarre
#CMD ["./main"]

FROM golang:alpine as builder

WORKDIR /app
ADD . /app

RUN go mod download
RUN go build -o serveur

FROM scratch 
COPY --from=builder /app/serveur serveur
CMD ["/serveur"]