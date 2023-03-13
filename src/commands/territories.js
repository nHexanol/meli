const { createCanvas, loadImage } = require('canvas');
require('discord.js');
const fs = require('fs');
require('canvas-webp');

// offset to map wynn coordinates to canvas coordinates
const offset = {
  x: 2382,
  z: 6537
};

module.exports = async function xp(guild, interaction) {

      await interaction.deferReply();
      const canvas = createCanvas(4034, 6414);
      const ctx = canvas.getContext('2d');

      // filter only guild territory, check from [terr].guild
      // [0] is territory name while [1] is territory data

      // load map
      await loadImage('../assets/map.png').then((image) => {
        ctx.drawImage(image, 0, 0, canvas.width, canvas.height);
        console.log('loaded map')
      });
      // draw trade routes
      // first, get the territories x and z coordinates so we know the center of the territory
      // so we can draw the trade route from the center of the territory to the center of the next territory

      await loadImage('../assets/route.png').then((image) => {
        ctx.globalAlpha = 0.67;
        ctx.drawImage(image, 0, 0, canvas.width, canvas.height);
        ctx.globalAlpha = 1;
        console.log('loaded routes')
      });
      // draw guild territories
      fetch('https://athena.wynntils.com/cache/get/territoryList')
        .then(r => r.json())
        .then(async function (t) {

          if (guild && guild.length >= 3 && guild.length <= 4 && ((!(/\d/gi).test(guild)) && (!guild.includes('_')))) {
            // test if its guild tag or not, if it is then get the guild name
            console.log('tag')
            // first find if guildtag is in the terr list or not, if not then the guild has no terrs
            var ts =  Object.entries(t);
            if (ts.find((a) => { a[1].guildPrefix === guild })?.length) {

              // guild has no territories
              console.log('no terrs');
              return await interaction.editReply({ content: `Guild '${guild}' has no territories.` });
            }
            else {
              // guild has territories
              // we will get the guild name from the guild tag from the guild list
              
              // guild name
              if (Object.entries(ts[0][1]).find(a => a[1].guildPrefix === guild) === undefined) {
                return interaction.editReply({ content: `Guild '${guild}' has no territories.` });
              }

              console.log(Object.entries(ts[0][1]).find(a => a[1].guildPrefix === guild)[1].guild)

              guild = Object.entries(ts[0][1]).find(a => a[1].guildPrefix === guild)[1].guild
            }

            
          }

          var terrs = Object.entries(t)


            // terrs[0] = terr name, terrs[1] = terr data
            // terrs[1].location.startX and startZ
            // terrs[1].location.endX and endZ
            console.log('\n');
            var realTerrs = Object.entries(terrs[0][1]);

            ctx.font = 'bold 40pt Ubuntu'
            ctx.textAlign = 'center';

            for (var i in realTerrs) {
              var locations = realTerrs[i][1].location;
              var startX = locations.startX + offset.x;
              var startZ = locations.startZ + offset.z;
              var endX = locations.endX + offset.x;
              var endZ = locations.endZ + offset.z;
              var centerX = (startX + endX) / 2;
              var centerZ = (startZ + endZ) / 2;

              var color = realTerrs[i][1].guildColor;
              var guildTag = realTerrs[i][1].guildPrefix;

              // draw the territory
              ctx.globalAlpha = 0.8;
              ctx.strokeStyle = color;

              // draw terrs outlines
              ctx.lineWidth = 3;
              ctx.strokeRect(startX, startZ, endX - startX, endZ - startZ);
              ctx.globalAlpha = 0.4;

              // if guild is provided then only colour the guild's territory
              // if not then just use the guild's default colour provided by the api
              if (!guild) {
                ctx.fillStyle = color;
              }
              else if (guild === realTerrs[i][1].guild && (guild !== undefined || guild !== null )) {
                ctx.fillStyle = color;
              }
              else if (guild !== realTerrs[i][1].guild && (guild !== undefined || guild !== null )) {
                // colour it grey
                ctx.fillStyle = '#444444';
              }

              ctx.lineWidth = 1;

              // then fill
              ctx.fillRect(startX, startZ, endX - startX, endZ - startZ);
              ctx.globalAlpha = 1;

              // draw guild tag but if the guild is provided then other ugild will be greyed out
              if (!guild) {
                ctx.globalAlpha = 1;
              }
              else if (guild === realTerrs[i][1].guild && (guild !== undefined || guild !== null )) {
                ctx.globalAlpha = 1;
              }
              else if (guild !== realTerrs[i][1].guild && (guild !== undefined || guild !== null )) {
                ctx.globalAlpha = 0.5;
              }

              ctx.fillStyle = '#ffffff';
              ctx.strokeStyle = '#000000';
              ctx.fillText(guildTag, centerX, centerZ + 20);
              ctx.strokeText(guildTag, centerX, centerZ + 20);

            }

            // we will create another canvas to crop the map to specific view/guild
            // then we will convert it to webp and send it to discord

            // first, we need to get all of the guild's territory coordinates
            // and find the min and max x and z coordinates

            // first, get all the indexes that contains terrs[i][1].guild === guild
            if (guild) {
            const guildTerritoriesListIndexes = [];
            const actualGuildTerritories = [];

            // get all the indexes of the guild's territories
            for (var i in realTerrs) {
              if (realTerrs[i][1].guild === guild) {
                guildTerritoriesListIndexes.push(i);
              }
            }

            // then get the real terrs object with the indexes
            guildTerritoriesListIndexes.map(function (i) {
              actualGuildTerritories.push(realTerrs[i][1]);
            })

            // find min and max x and z coordinates
            // we will use these to crop the map

            // initialize min and max x and z coordinates lol
            var minX = 9999;
            var maxX = -9999;
            var minZ = 9999;
            var maxZ = -9999;

            // goofy ah way to get the min and max
            for (var i in actualGuildTerritories) {
              var locations = actualGuildTerritories[i].location;
              var startX = locations.startX + offset.x;
              if (startX < minX) minX = startX;
              if (startX > maxX) maxX = startX;
              var startZ = locations.startZ + offset.z;
              if (startZ < minZ) minZ = startZ;
              if (startZ > maxZ) maxZ = startZ;
            }

            // now we have the min and max x and z coordinates we can crop the map
            // we will create a canvas with the size of the cropped map plus some padding
            // then we will draw the cropped map on the canvas
            console.log(`Size of the canvas will be ${maxX - minX + 300} x ${maxZ - minZ + 300} pixels.`)
            const canvasGuild = createCanvas(maxX - minX + 300, maxZ - minZ + 300);
            const ctxGuild = canvasGuild.getContext('2d');

            // draw the cropped map
            ctxGuild.drawImage(canvas, minX - 150, minZ - 150, maxX - minX + 300, maxZ - minZ + 300, 0, 0, maxX - minX + 300, maxZ - minZ + 300);
            fs.writeFileSync('./out.webp', canvasGuild.toBuffer('image/webp', { lossless: false, quality: 0.8 }));
          }
          else {
            // draw the full map
            canvas.toBuffer('image/webp', { lossless: false, quality: 0.8 });
            fs.writeFileSync('./out.webp', canvas.toBuffer('image/webp', { lossless: false, quality: 0.8 }));
          }
            // write to file first and attach the file or else we will get AbortError when sending the file

          await interaction.editReply({ files: [ './out.webp' ] });

        });
}