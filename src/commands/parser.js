const Discord = require('discord.js');
require('discord.js');
const xp = require('./xp.js');
const territories = require('./territories.js');
const getGuildFromTag = require('../utils/tagApiResponse.js');

module.exports = async function (interaction) {
  console.log('parser called')
  // for guild 
  if (interaction.customId.includes('onlinemember_')) {
    // query guild members
    var guildName = interaction.customId.replace(/-/gi, ' ').split('_')[1];
    // fetch later

  }
  else if (interaction.customId.includes('xp_')) {
    // query top 5 xp
    var guildNameFull = interaction.customId.split('_')[1].replace(/-/gi, ' ');
    xp(guildNameFull, interaction);

  }
  else if (interaction.customId.includes('activities_')) {
    // query top 5 activities later, need a proper db

  }
  else if (interaction.customId.includes('territories')) {
    var guildNameFull = interaction.customId.split('_')[1].replace(/-/gi, ' ');
      territories(guildNameFull, interaction);
  }
}