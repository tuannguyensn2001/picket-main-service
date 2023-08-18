import http from 'k6/http';

export const options = {
    vus: 100,
    duration: '30s'
}

export default function() {
    const url = "http://localhost:21000/api/v1/answersheets/start";
    // const url = "http://localhost:21000/health"
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2ODMzNTczOTMsIm5iZiI6MTY4MzI3MDk5MywiaWF0IjoxNjgzMjcwOTkzLCJqdGkiOiIxIn0.IprqXjlfo4MSnG3TXABwuJe5fStAuf32008bK3LxCvo";

    const params = {
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
        },
    }
    const payload = JSON.stringify({
        test_id: 1,
        question_id: 1,
        answer: "A"
    })
    // console.log(payload)

    http.post(url,payload , params)
}