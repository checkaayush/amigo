#########################################
## Multi Stage Docker build (Recommended)
#########################################

FROM golang:1.11 AS builder

# Download and install the latest release of dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/github.com/checkaayush/amigo
COPY Gopkg.toml Gopkg.lock ./

# Ensure dependencies
RUN dep ensure --vendor-only

# Copy the code from the host and compile it
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE 5000
CMD [ "./app" ]

###########################################


#####################################################
## Multi Stage Docker build with private dependencies
## Refer: https://bit.ly/2oY3pCn
#####################################################

# FROM golang:1.11 AS builder

# # Download and install the latest release of dep
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# WORKDIR $GOPATH/src/github.com/checkaayush/amigo
# COPY Gopkg.toml Gopkg.lock ./

# # Ensure dependencies
# ARG SSH_PRIVATE_KEY
# RUN eval "$(ssh-agent -s)" && \
#     mkdir -p /root/.ssh && \
#     echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && \
#     chmod 600 /root/.ssh/id_rsa && \
#     ssh-add /root/.ssh/id_rsa && \
#     ssh-keyscan github.com > /root/.ssh/known_hosts && \
#     dep ensure --vendor-only && \
#     rm /root/.ssh/id_rsa && \
#     rm -rf /root/.ssh/

# # Copy the code from the host and compile it
# COPY . ./
# RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

# FROM scratch
# COPY --from=builder /app ./
# EXPOSE 5000
# CMD [ "./app" ]

###########################################


############################
## Single Stage Docker build
############################

# FROM golang:1.11

# # Download and install the latest release of dep
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# WORKDIR $GOPATH/src/github.com/checkaayush/amigo
# COPY Gopkg.toml Gopkg.lock ./

# # Ensure dependencies
# RUN dep ensure --vendor-only

# # Copy the code from the host and compile it
# COPY . ./
# RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

# EXPOSE 5000
# CMD [ "/app" ]

#############################


######################################################
## Single Stage Docker build with private dependencies
## Less secure; not recommended. Refer: https://bit.ly/2oY3pCn
######################################################

# FROM golang:1.11

# # Download and install the latest release of dep
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# WORKDIR $GOPATH/src/github.com/checkaayush/amigo
# COPY Gopkg.toml Gopkg.lock ./

# # Ensure dependencies
# ARG SSH_PRIVATE_KEY
# RUN eval "$(ssh-agent -s)" && \
#     mkdir -p /root/.ssh && \
#     echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && \
#     chmod 600 /root/.ssh/id_rsa && \
#     ssh-add /root/.ssh/id_rsa && \
#     ssh-keyscan github.com > /root/.ssh/known_hosts && \
#     dep ensure --vendor-only && \
#     rm /root/.ssh/id_rsa && \
#     rm -rf /root/.ssh/

# # Copy the code from the host and compile it
# COPY . ./
# RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

# EXPOSE 5000
# CMD [ "/app" ]

########################################################
