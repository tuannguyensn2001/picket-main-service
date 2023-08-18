import http from 'k6/http';

export const options = {
    vus: 100,
    duration: '30s'
}

export default function () {
    const url = "http://localhost:21000/api/v1/answersheets/answer";
    // const url = "http://localhost:21000/health"
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2ODU0MTY0MzgsIm5iZiI6MTY4NTMzMDAzOCwiaWF0IjoxNjg1MzMwMDM4LCJqdGkiOiIxIn0.ABTAyz2-nFqUfIJyAWgr1AfLFk2n_kNtYgVFC4oQIE4";

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