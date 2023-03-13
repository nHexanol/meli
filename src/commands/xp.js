const { createCanvas, loadImage } = require('canvas');
require('discord.js');

module.exports = function xp(guild, interaction) {
  console.log('called');
  console.log(guild)
  fetch(`https://api.wynncraft.com/public_api.php?action=guildStats&command=${guild}`)
    .then(r => r.json())
    .then(async function (json) {
      await interaction.deferReply();
      const canvas = createCanvas(1600, 900);
      const ctx = canvas.getContext('2d');

      await loadImage('../assets/0.png').then((image) => {
        ctx.drawImage(image, 0, 0, canvas.width, canvas.height);
      });

      ctx.font = 'bold 55pt Ubuntu';
      ctx.textAlign = 'left';
      ctx.fillStyle = '#fff';
      ctx.fillText(`${json.name}`, 75, 125);
      var length = ctx.measureText(json.name).width;
      ctx.font = '55pt Ubuntu';
      ctx.fillText(`XP Contribution`, 125 + length, 125);

      // space by 75

      var members = json.members;
      console.log(json);
      members.sort(function (a, b) {
        // from highest to lowest
        return b.contributed - a.contributed;
      });
      var top8 = members.slice(0, 8);
      for (var i in top8) {
        // start from y = 300
        // name rank xp
        ctx.font = 'bold 35pt Ubuntu';
        ctx.textAlign = 'left';
        ctx.fillText(`[ ${parseInt(i) + 1} ] ${top8[i].name}`, 75, 250 + (i * 75));

        ctx.font = '35pt Ubuntu';
        ctx.textAlign = 'center';
        ctx.fillText(`${top8[i].rank}`, 725, 250 + (i * 75));

        ctx.font = '35pt Ubuntu Mono';
        ctx.textAlign = 'right';
        ctx.fillText(`${top8[i].contributed.toLocaleString('en-US')}`, 1525 - 75, 250 + (i * 75));

      }

      await interaction.editReply({ files: [canvas.toBuffer()] });

    });
}