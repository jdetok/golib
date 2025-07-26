# GETENV GO PACKAGE
- derived from getenv in go-api-jdeko.me
basic package using github.com/joho/godotenv to read and return variables from .env file
# how to use
- LoadDotEnv must be called first, preferably only once at the beginning of the runtime
    - this reads the .env file and loads its contents as environment vars
- after loading, GetEnvStr(VARNAME) & GetEnvInt(VARNAME) can be called to return strings or ints from env, respectively