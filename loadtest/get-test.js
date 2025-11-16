import http from 'k6/http';
import { sleep, check } from 'k6';
import { Rate, Trend } from 'k6/metrics';

export let responseTime = new Trend('response_time');
export let successRate = new Rate('success_rate');

export let options = {
    vus: 10,
    duration: '1m',
    rps: 5,
    thresholds: {
        'response_time': ['p(95)<300'],
        'success_rate': ['rate>0.999'],
    },
};

const teams = Array.from({ length: 20 }, (_, i) => `Team${i+1}`);
const users = Array.from({ length: 200 }, (_, i) => `user${i+1}`);

const BASE_URL = 'http://localhost:8080';

// GET /pullRequest/stats
function getPullRequestStats(team) {
    let url = `${BASE_URL}/pullRequest/stats?team_name=${team}`;
    let res = http.get(url);
    responseTime.add(res.timings.duration);
    successRate.add(res.status === 200 || res.status === 404);
}

// GET /team/get
function getTeam(team) {
    let url = `${BASE_URL}/team/get?team_name=${team}`;
    let res = http.get(url);
    responseTime.add(res.timings.duration);
    successRate.add(res.status === 200 || res.status === 404);
}

// GET /users/getReview
function getUserAssignedPullRequests(user) {
    let url = `${BASE_URL}/users/getReview?user_id=${user}`;
    let res = http.get(url);
    responseTime.add(res.timings.duration);
    successRate.add(res.status === 200 || res.status === 404);
}

export default function () {
    const team = teams[Math.floor(Math.random() * teams.length)];
    const user = users[Math.floor(Math.random() * users.length)];

    getPullRequestStats(team);
    getTeam(team);
    getUserAssignedPullRequests(user);

    sleep(0.2);
}
