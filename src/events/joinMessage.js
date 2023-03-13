const { createCanvas, loadImage } = require('canvas');
const Discord = require('discord.js');

module.exports = async function a(member) {
    const channel = member.guild.channels.cache.find(ch => ch.name === 'general');
    if (!channel) return 1;

    const canvas = createCanvas(700, 250);
    const ctx = canvas.getContext('2d');

    loadImage('./images/welcome.png').then(async (image) => {
        ctx.drawImage(image, 0, 0, canvas.width, canvas.height);

        ctx.strokeStyle = '#74037b';
        ctx.strokeRect(0, 0, canvas.width, canvas.height);

        ctx.font = '28px sans-serif';
        ctx.fillStyle = '#ffffff';
        ctx.fillText('Welcome to the server,', canvas.width / 2.5, canvas.height / 3.5);

        ctx.font = applyText(canvas, `${member.displayName}!`);
        ctx.fillStyle = '#ffffff';
        ctx.fillText(`${member.displayName}!`, canvas.width / 2.5, canvas.height / 1.8);

        ctx.beginPath();
        ctx.arc(125, 125, 100, 0, Math.PI * 2, true);
        ctx.closePath();
        ctx.clip();

        const avatar = await loadImage(member.user.displayAvatarURL({ format: 'jpg' }));
        ctx.drawImage(avatar, 25, 25, 200, 200);

        const attachment = new Discord.MessageAttachment(canvas.toBuffer(), 'welcome-image.png');

        channel.send({ files : [ attachment ] });
    });
}