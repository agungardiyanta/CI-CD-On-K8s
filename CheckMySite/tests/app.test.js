import { expect } from 'chai';
import sinon from 'sinon';
import fetch from 'node-fetch';
import request from 'supertest';
import app from '../app.js'; // Import the Express app

describe('GET site/:domain', function() {
    let fetchStub;

    // Before each test, stub the fetch function
    beforeEach(() => {
        fetchStub = sinon.stub(fetch, 'default');
    });

    // After each test, restore the fetch function
    afterEach(() => {
        fetchStub.restore();
    });

    it('should render status.ejs with isUp=true for a site that is up', async function() {
        // Arrange: Stub fetch to return a successful response
        fetchStub.resolves({
            ok: true,
            status: 200,
        });

        // Act: Make a request to the app
        const res = await request(app).get('/example.com');

        // Assert: Check that the response contains the expected content
        expect(res.status).to.equal(200);
        expect(res.text).to.include('example.com');
        expect(res.text).to.include('class="status-up"');
    });

    it('should render status.ejs with isUp=false for a site that is down', async function() {
        // Arrange: Stub fetch to return a failed response
        fetchStub.resolves({
            ok: false,
            status: 500,
        });

        // Act: Make a request to the app
        const res = await request(app).get('/ybe.com');

        // Assert: Check that the response contains the expected content
        expect(res.status).to.equal(200);
        expect(res.text).to.include('ybe.com');
        expect(res.text).to.include('class="status-down"');
    });

    it('should render status.ejs with isUp=false when fetch throws an error', async function() {
        // Arrange: Stub fetch to throw an error
        fetchStub.rejects(new Error('Network error'));

        // Act: Make a request to the app
        const res = await request(app).get('/ybe.com');

        // Assert: Check that the response contains the expected content
        expect(res.status).to.equal(200);
        expect(res.text).to.include('ybe.com');
        expect(res.text).to.include('class="status-down"');
    });
});
