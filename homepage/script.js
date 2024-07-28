const soundscape = document.getElementById('soundscape');  

// 随机生成声音节点  
const sounds = [  
    { name: '声音1', url: 'sound1.mp3' },  
    { name: '声音2', url: 'sound2.mp3' },  
    { name: '声音3', url: 'sound3.mp3' },  
    { name: '声音4', url: 'sound4.mp3' },  
    { name: '声音5', url: 'sound5.mp3' }  
];  

// 随机打乱声音数组  
const shuffleSounds = () => {  
    for (let i = sounds.length - 1; i > 0; i--) {  
        const j = Math.floor(Math.random() * (i + 1));  
        [sounds[i], sounds[j]] = [sounds[j], sounds[i]];  
    }  
};  

// 创建声音节点  
const createSoundNodes = () => {  
    shuffleSounds();  
    sounds.forEach((sound) => {  
        const node = document.createElement('div');  
        node.classList.add('sound-node');  
        node.innerText = sound.name;  

        // 点击事件  
        node.addEventListener('click', () => {  
            playSound(sound.url);  
        });  

        // 长按事件  
        let pressTimer;  
        node.addEventListener('mousedown', () => {  
            pressTimer = setTimeout(() => {  
                alert(`你长按了${sound.name}`);  
            }, 700);  
        });  
        node.addEventListener('mouseup', () => {  
            clearTimeout(pressTimer);  
        });  
        node.addEventListener('mouseleave', () => {  
            clearTimeout(pressTimer);  
        });  

        soundscape.appendChild(node);  
    });  
};  

// 播放声音  
const playSound = (url) => {  
    const audio = new Audio(url);  
    audio.play();  
};  

// 初始化声音节点  
createSoundNodes();