#!/bin/bash

SECRET_FOLDER="/secrets"
TRUSTSTORE_FILE="truststore.jks"
TRUSTSTORE12_FILE="truststore.p12"

#if [ -f $SECRET_FOLDER/$TRUSTSTORE12_FILE ]; then
#    rm -f $($SECRET_FOLDER/$TRUSTSTORE12_FILE)
#fi

#if [ ! -f $SECRET_FOLDER/$TRUSTSTORE_FILE ]; then
#    echo "Файл '$SECRET_FOLDER/$TRUSTSTORE_FILE' не существует."
#
#    exit 1
#fi

export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

keytool -importkeystore -srckeystore $SECRET_FOLDER/$TRUSTSTORE_FILE -destkeystore $SECRET_FOLDER/$TRUSTSTORE12_FILE \ 
-srcstoretype JKS -deststoretype PKCS12 -srcstorepass $GO_PHDOCBASEDBBZ_KSSLPASSWORD -deststorepass $GO_PHDOCBASEDBBZ_KSSLPASSWORD