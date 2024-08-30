import { Client } from 'pg';

export default async function insertEventData(userId, eventName, eventData) {
    const client = new Client({
        user: process.env.DB_USER,
        host: 'db',
        database: process.env.DB_DATABASE,
        password: process.env.DB_PASSWORD,
        port: process.env.DB_PORT,
    });
    await client.connect();

    await client.query(`
    INSERT INTO events (user_id, event_name, event_data)
    VALUES ($1, $2, $3)
  `, [userId, eventName, eventData]);

    await client.end();
};