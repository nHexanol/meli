require('discord.js');
const { createCanvas, loadImage } = require('canvas');
const guildCanvas = require('../utils/guildCanvas');
const charactersButton = require('./characters.js');

module.exports = async function (interaction) {

  if (interaction.customId.includes('characters+')) {
    var username = interaction.customId.split('+')[1];
    var classUuid = interaction.customId.split('+')[2];

    charactersButton(username, classUuid, interaction);


  }

	if (interaction.customId === 'guild') {
		var guildName = interaction.values[0];
		queryGuild(guildName).then(function (g) {
			// todo : canvas shit
			guildCanvas(g, interaction);
		});
	}

	

	function queryGuild(guildName) {
		return fetch(`https://api.wynncraft.com/public_api.php?action=guildStats&command=${guildName}`)
			.then(r => r.json())
			.then(function (json) {
				if (json.error) {
					return;
				}
				else {
					console.log(json);
					return Promise.resolve(json);
				}
			});
	}

}