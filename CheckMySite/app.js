import express from 'express';
import fetch from 'node-fetch';
import path from 'path';
import { fileURLToPath } from 'url';

const app = express();

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Set the view engine to use EJS templates
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'templates'));

// Function to check if the site is up
async function checkSite(url) {
    try {
        const response = await fetch(url, { method: 'HEAD', timeout: 10000 });
        return response.ok;
    } catch (error) {
        return false;
    }
}

// Route to check the site status
app.get('/site/:domain', async (req, res) => {
    const domain = req.params.domain;
    const url = `http://${domain}`;

    const isUp = await checkSite(url);
    if (isUp){
        const status = true;
        res.render('status', { domain, status });
    }
    else{
        const status = false;
        res.render('status', { domain, status });
    }
    
});

// Start the server
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});

export default app;