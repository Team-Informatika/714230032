// Smooth scroll for nav links
document.querySelectorAll('.nav-links a').forEach(link => {
    link.addEventListener('click', e => {
      e.preventDefault();
      const targetId = link.getAttribute('href').replace('#', '');
      const targetEl = document.getElementById(targetId);
      targetEl.scrollIntoView({ behavior: 'smooth' });
    });
  });
  
  // Button interaction
  document.querySelector('.btn-primary')?.addEventListener('click', () => {
    document.getElementById('contact').scrollIntoView({ behavior: 'smooth' });
  });
  

const text = "Halo, Saya Ode Andi Alamsyah";
let i = 0;
const typedText = document.getElementById("typed-text");

function type() {
    if (i < text.length) {
        typedText.innerHTML += text.charAt(i);
        i++;
        setTimeout(type, 100);
    }
}

document.addEventListener("DOMContentLoaded", type);  