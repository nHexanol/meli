const Discord = require('discord.js');
const jsonPreprocessor = require('../utils/userJsonPreprocessor.js');
const getGuildFromTag = require('../utils/tagApiResponse.js');
const guildCanvas = require('../utils/guildCanvas.js');
const playerCanvas = require('../utils/playerCanvas.js');

const rows = [];

module.exports = async function (interaction) {

  var input = interaction.options.getString('query')
  console.log(input);
  if (input === null) {
    // no input means to get the discord id, then find it in the userData.json and get the uuid
    var userId = interaction.user.id
    var users = JSON.parse(jsonPreprocessor())
    var uuid = Object.entries(users).find(el => el[1] == userId)[0]
    return await queryPlayer(uuid);
  }

  if (input.length >= 3 && input.length <= 4 && ((!(/\d/gi).test(input)) && (!input.includes('_')))) {
    // input is a guild tag
      console.log('GUILD TAG');
      var resp = await getGuildFromTag(input);
      if (resp.type === 'GUILD') {
        var guildName = resp.guild;
        console.log(resp.guild);
        var guild = await queryGuild(guildName);
        await guildCanvas(guild, interaction);
      }
      else if (resp.type === 'MULTIPLE_GUILDS') {
        // make a prompt to select a guild
        console.log('MULTIPLE GUILD')

        var index = 0;
        const options = [];

        for (const [key, value] of Object.entries(resp.guilds)) {
          // make a button for each guild when clicked
          console.log(key, value);
          options[index] = {
            label: `[${key}] ${value}`,
            description: `Query stats for guild '${value}'`,
            value: value,
          };
          index++;
        }
        console.log(options);

        const rows = new Discord.ActionRowBuilder()
          .addComponents(
            new Discord.StringSelectMenuBuilder()
              .setCustomId(`guild`)
              .setPlaceholder('Select a guild')
              .addOptions(options)
          )
        interaction.reply({ content: `Multiple guild with provided tag.`, components: [rows] });
      }
  }
  else {
    console.log('not tag');
    try {
      var player = await queryPlayer(input);
      var guild = await queryGuild(input);

      console.log(player)

      if (player && guild) {
        // both player and guild exist
        // make a prompt to select a guild

        const options = [{
          label: `Player: ${player.data[0].username}`,
          description: `Query stats for player '${player.data[0].username}'`,
          value: player.data[0].username,
        },
        {
          label: `Guild: ${guild.name}`,
          description: `Query stats for guild '${guild.name}'`,
          value: guild.name,
        }];

        const rows = new Discord.ActionRowBuilder()
          .addComponents(
            new Discord.StringSelectMenuBuilder()
              .setCustomId(`guild`)
              .setPlaceholder('Select a guild')
              .addOptions(options)
          )
        interaction.reply({ content: 'Multiple resolves received.', components: [rows] });

      }
      else if (!guild) {
        // player
        console.log('PLAYER');
        console.log(interaction);
        playerCanvas(player, interaction);
        // TODO: make a player canvas
      }
      else if (!player) {
        // guild
        console.log('GUILD');
        guildCanvas(guild, interaction);
        // TODO: make a guild canvas with an option to select a player
      }
      else {
        // no player or guild
        interaction.reply('No player or guild found.');
      }
    }
    catch (error) {

    }

  }

  async function queryPlayer(username) {
    console.log('getting player')
    const r = await fetch(`https://api.wynncraft.com/v2/player/${username}/stats`);
    const u = await r.json();
    if (u.code === 400) {
      return;
    }
    else {
      console.log('got player');
      playerCanvas(u, interaction);
      return json;
    }
  }

  async function queryGuild(guildName) {
    console.log('getting guild')
    return fetch(`https://api.wynncraft.com/public_api.php?action=guildStats&command=${guildName}`)
      .then(r => r.json())
      .then(function (json) {
        if (json.error) {
          return;
        }
        else {
          console.log('got guild')
          return json;
        }
      });
  }
}
