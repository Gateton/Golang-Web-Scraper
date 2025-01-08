const axios = require('axios');
const cheerio = require('cheerio');
const fs = require('fs');
const yargs = require('yargs/yargs');
const { hideBin } = require('yargs/helpers');

const argv = yargs(hideBin(process.argv))
  .option('url', {
    alias: 'u',
    description: 'URL to scrape',
    type: 'string',
    demandOption: true
  })
  .option('output', {
    alias: 'o',
    description: 'Output JSON file',
    type: 'string',
    default: 'output.json'
  })
  .option('header', {
    alias: 'H',
    description: 'Custom headers',
    type: 'array'
  })
  .option('cookie', {
    alias: 'C',
    description: 'Custom cookies',
    type: 'array'
  })
  .argv;

async function scrape(url, headers, cookies) {
  try {
    const config = { headers: {} };
    
    if (headers) {
      headers.forEach(header => {
        const [key, value] = header.split(':');
        config.headers[key.trim()] = value.trim();
      });
    }

    if (cookies) {
      config.headers.Cookie = cookies.join('; ');
    }

    const response = await axios.get(url, config);
    const $ = cheerio.load(response.data);

    const data = {
      titles: [],
      prices: [],
      images: []
    };

    // Scrape titles
    $('h2.product-title').each((i, el) => {
      data.titles.push($(el).text().trim());
    });

    // Scrape prices
    $('span.price').each((i, el) => {
      data.prices.push($(el).text().trim());
    });

    // Scrape image URLs
    $('img.product-image').each((i, el) => {
      data.images.push($(el).attr('src'));
    });

    return data;
  } catch (error) {
    console.error(`Error scraping '${url}':`, error.message);
    process.exit(1);
  }
}

async function main() {
  const { url, output, header, cookie } = argv;

  const data = await scrape(url, header, cookie);

  fs.writeFileSync(output, JSON.stringify(data, null, 2));
  console.log(`Scraping completed. Data saved to ${output}`);
}

main();
