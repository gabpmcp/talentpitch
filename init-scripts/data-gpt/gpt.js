import { post } from 'axios';

export default async function fetchGPTData() {
    const response = await post('https://api.openai.com/v1/chat/completions', {
        headers: {
            'Authorization': `Bearer your_api_key`,
            'Content-Type': 'application/json'
        },
        data: {
            model: "gpt-4",
            messages: [{ role: "system", content: "Your prompt here" }]
        }
    });
    return response.data;
};
