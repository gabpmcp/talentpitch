import fetchGPTData from './api/gpt';
import insertEventData from './db/insert';

(async function main() {
    const userId = 'your-uuid';
    const eventName = 'GPT Response';
    const eventData = await fetchGPTData();

    await insertEventData(userId, eventName, eventData);
    console.log('Process completed successfully');
})();
