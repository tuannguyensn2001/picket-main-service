import http from 'k6/http';

export const options = {
    vus: 100,
    duration: '30s'
}

export default function () {
    const url = "http://localhost:21000/api/v1/answersheets/submit";
    // const url = "http://localhost:21000/health"
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2ODEyMjk1MjAsIm5iZiI6MTY4MTE0MzEyMCwiaWF0IjoxNjgxMTQzMTIwLCJqdGkiOiIxIn0.XBP7NxskcGjRtkD2W_sAk_0Yo-fJBbVPWg4G527rHXg";

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
    console.log(payload)

    http.post(url,payload , params)
}