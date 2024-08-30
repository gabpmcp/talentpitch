# Utiliza la última versión LTS de Node.js como base
FROM node:lts

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /usr/src/app

# Copia el package.json y package-lock.json (si existe)
COPY package*.json ./

# Instala las dependencias de la aplicación
RUN npm install --production

# Copia el resto de la aplicación al contenedor
COPY . .

# Expone el puerto en el que la aplicación va a correr (si es necesario)
# EXPOSE 3000

# Define la variable de entorno para el script de inicialización
ENV NODE_ENV=production

# Comando para ejecutar el script que poblará la tabla de eventos
CMD ["node", "index.js"]
