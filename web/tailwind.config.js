module.exports = {
  purge: false,
  // darkMode: false, // or 'media' or 'class'
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      height: {
        '500-px': '500px',
      },
      maxWidth: {
        '150-px': '150px',
        '1/4': '25%',
        '1/2': '50%',
        '3/4': '75%',
      },
      borderColor: {
        'blueGray-200': 'rgba(226, 232, 240, 1)',
      },
    }, // 进行继承并补充 而不是完全覆盖
    minHeight: {
      '0': '0',
      '1/5': '20%',
      '1/2': '50%',
      '3/4': '75%',
      'full': '100%',
    },
    zIndex: {
      0: 0,
      10: 10,
      20: 20,
      30: 30,
      40: 40,
      50: 50,
      25: 25,
      75: 75,
      99: 99,
      100: 100,
      auto: 'auto',
    },
  },
  variants: {},
  plugins: [],
}
