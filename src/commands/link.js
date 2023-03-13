require('discord.js');
const { exec } = require('child_process');

module.exports = async function (interaction) {
    var discord = interaction.options.getUser('discord');
    if (discord === null) {
        var discordUser = interaction.user.id;
    }
    else {
        var discordUser = discord.id;
    }
    console.log(discordUser);
    var username = interaction.options.getString('username');

    fetch(`https://api.wynncraft.com/v2/player/${username}/stats`)
    .then(r => r.json())
    .then(function (json) {
        var uuid = json.data[0].uuid;
        const linker = exec(`../bin/cacheUserData ${uuid} ${discordUser}`);
        console.log(uuid)
        linker.once('exit', function (code) {
            if (code === 0) {
                interaction.reply('Link successful');
            }
            else {
                interaction.reply('Link error');
            }
        });
    })
    .catch(function (error) {
        interaction.reply('Link error');
    })
}