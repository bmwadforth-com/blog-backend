import React, {useEffect, useState} from 'react';
import {QueryRequest} from '../../protocol_buffers/gemini_service_pb';
import {GeminiClient} from '../../protocol_buffers/Gemini_serviceServiceClientPb';

const Gemini = () => {
    const [responseMessages, setResponseMessages] = useState([]);
    const [query, setQuery] = useState('');

    useEffect(() => {
        if (query !== '') {
            const client = new GeminiClient('http://localhost:8081', null, null);

            const request = new QueryRequest()
            request.setQuery(query);
            const stream = client.queryGemini(request);

            stream.on('data', (response: any) => {
                console.log(response);
                //setResponseMessages((prevMessages) => [...prevMessages, response.getMessage()]);
            });

            stream.on('status', (status: any) => {
                if (status.code !== 0) {
                    console.log('Error code:', status.code, 'Error message:', status.details);
                }
            });

            stream.on('end', () => {
                console.log('Stream ended.');
            });
        }
    }, [query]);

    const handleQueryChange = (event: any) => {
        setQuery(event.target.value);
    };

    const handleSubmit = (event: any) => {
        event.preventDefault();
        setResponseMessages([]); // Clear previous messages
        // You might want to add logic to trigger the useEffect
    };

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <input type="text" value={query} onChange={handleQueryChange}/>
                <button type="submit">Send Query</button>
            </form>
            <div>
                <h2>Responses:</h2>
                {responseMessages.map((msg, index) => (
                    <p key={index}>{msg}</p>
                ))}
            </div>
        </div>
    );
};

export default Gemini;
