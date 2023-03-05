import http from 'k6/http';

export const options = {
    vus: 200,
    duration: '2s'
}

export default function() {
    const url = "http://localhost:21000/api/v1/answersheets/test/1/content";
    // const url = "http://localhost:21000/health"
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2NzgwMDk2MTUsIm5iZiI6MTY3NzkyMzIxNSwiaWF0IjoxNjc3OTIzMjE1LCJqdGkiOiIxIn0.s7opmBujchs4QZ48cnTZx59Rwl24U5dsN_mUVCcDfJ0";

    const params = {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    }

    http.get(url,params)
}