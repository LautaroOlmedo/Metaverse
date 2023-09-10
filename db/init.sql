-- CREATE DATABASE IF NOT EXISTS lautaroolmedo
SELECT 'CREATE DATABASE lautaroolmedo'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'metaverse-db')\gexec