const audioCtx = new AudioContext();

const wave = new PeriodicWave(audioCtx, {
    real: wavetable.real,
    imag: wavetable.imag,
  });

  function playSweep(time) {
    const osc = new OscillatorNode(audioCtx, {
      frequency: 380,
      type: "custom",
      periodicWave: wave,
    });
    osc.connect(audioCtx.destination);
    osc.start(time);
    osc.stop(time + 1);
  }

  