module.exports = async function getGuildFromTag(tag) {
    const url = `http://avicia.ga/api/tag/?tag=${tag}`;
    const headers = {
      Accept: 'application/json'
    };
      const response = await fetch(url, { headers });
      if (!response.ok) {
        throw new Error(`Failed to fetch data: ${response.status} ${response.statusText}`);
      }
  
      const body = await response.text();
      const json = JSON.parse(body);
  
      if (typeof json === 'object') {
        const guilds = {};
        for (const key in json) {
          if (Object.prototype.hasOwnProperty.call(json, key) && typeof json[key] === 'string') {
            guilds[key] = json[key];
          }
        }
        console.log(guilds)
        return { type: 'MULTIPLE_GUILDS', guilds };
      } else if (typeof json === 'string') {
        const guild = json.trim();
        if (guild === 'null') {
          throw new Error(`No guild found with tag ${tag}`);
        }
        return { type: 'GUILD', guild };
      } else {
        throw new Error(`Invalid response type: ${typeof json}`);
      }
  }
  