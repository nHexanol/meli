const Discord = require('discord.js');
const { createCanvas, loadImage } = require('canvas');

module.exports = async function (username, classUuid, interaction) {
  await interaction.deferReply();
  fetch(`https://api.wynncraft.com/v2/player/${username}/stats`)
  .then(res => res.json())
  .then(async function (c) {
    c = c.data[0].characters[classUuid];
    console.log(c)
    console.log(c.professions)
    // split into 3 canvas
    // first is prof
    // second is raids
    // third is dungeons

    const profCanvas = createCanvas(1600, 900);
    const ctx = profCanvas.getContext('2d');
    await loadImage('../assets/0.png').then((image) => { ctx.drawImage(image, 0, 0, 1920, 1080); });

    ctx.font = 'bold 55pt Ubuntu';
    ctx.textAlign = 'left';
    ctx.fillStyle = '#ffffff';
    ctx.strokeStyle = '#000000';
    ctx.fillText(`Professions`, 75, 125);
    ctx.strokeText(`Professions`, 75, 125);
    var textWidth = ctx.measureText('Professions').width;

    // start at 75, 300
    ctx.font = 'bold 35pt Ubuntu';
    ctx.fillText(`Mining : `, 75, 200);
    var textWidth = ctx.measureText('Mining : ').width;
    ctx.font = '35pt Ubuntu';
    ctx.fillText(`Lv. ${c.professions.mining.level} [ ${c.professions.mining.xp} % ]`, 75 + textWidth, 200);

    ctx.font = 'bold 35pt Ubuntu';
    ctx.fillText(`Woodcutting : `, 75, 260);
    var textWidth = ctx.measureText('Woodcutting : ').width;
    ctx.font = '35pt Ubuntu';
    ctx.fillText(`Lv.  ${c.professions.woodcutting.level} [ ${c.professions.woodcutting.xp} % ]`, 75 + textWidth, 260);

    ctx.font = 'bold 35pt Ubuntu';
    ctx.fillText(`Farming : `, 75, 320);
    var textWidth = ctx.measureText('Farming : ').width;
    ctx.font = '35pt Ubuntu';
    ctx.fillText(`Lv.  ${c.professions.farming.level} [ ${c.professions.farming.xp} % ]`, 75 + textWidth, 320);

    ctx.font = 'bold 35pt Ubuntu';
    ctx.fillText(`Fishing : `, 75, 380);
    var textWidth = ctx.measureText('Fishing : ').width;
    ctx.font = '35pt Ubuntu';
    ctx.fillText(`Lv.  ${c.professions.fishing.level} [ ${c.professions.fishing.xp} % ]`, 75 + textWidth, 380);

    await interaction.editReply({ files: [ profCanvas.toBuffer() ] });

  })
  .catch(function (e) {
    interaction.editReply({ content: `Error: ${e}` });
  })
}