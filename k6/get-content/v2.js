import http from 'k6/http';

export const options = {
    vus: 100,
    duration: '30s'
}

export default function() {
    const url = "http://localhost:21000/api/v2/answersheets/test/1/content";
    // const url = "http://localhost:21000/health"

    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2ODM1MjU0OTYsIm5iZiI6MTY4MzQzOTA5NiwiaWF0IjoxNjgzNDM5MDk2LCJqdGkiOiIxIn0.0VebaO0ItHet5ZsKUg0FstL4BBwhTYiFjMplwYQGjDk"
    const params = {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    }

    http.get(url,params)
}