FROM ubuntu

#Dépendances 
RUN apt update && apt install python build-essential gcc g++ git

RUN git clone https://github.com/nodejs/node.git nodejs

WORKDIR nodejs


#Debut Compiltation
#Vérifie que tout est bien présent lors de la compilation
RUN ./configure  
RUN make -j 4
RUN make install
#Fin Compiltation

#Commande nécessaire pour l'utilisation 
ENTRYPOINT [ "/usr/local/bin/node" ]
CMD ["-e", "console.('hello 'wo, process.version)"]