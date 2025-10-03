/**
 * MarkovChain is a text generator program.
 *
 * It takes in a seed text to model markov chains with probability paths.
 * To generate text, it takes in a starting word and number of words to generate.
 */
class MarkovChain {
  #model = {};

  modelChain(seedTxt) {
    const words = seedTxt.split(" ");
    for (let i = 0; i < words.length - 2; i++) {
      const currWord = words[i],
        nextWord = words[i + 1];

      if (!this.#model[currWord]) {
        this.#model[currWord] = {
          [nextWord]: 1,
        };
        continue;
      }

      if (!this.#model[currWord][nextWord]) {
        this.#model[currWord][nextWord] = 1;
        continue;
      }

      this.#model[currWord][nextWord]++;
    }

    // console.log(this.#model);
  }

  generateText(startingWord, wordLimit) {
    if (!this.#model[startingWord]) {
      return "";
    }

    let nextWord = startingWord;
    let curLimit = 1;

    while (curLimit < wordLimit) {
      console.log(nextWord + " ");

      const node = this.#model[nextWord];
      const sumCnt = Object.values(node).reduce((s, c) => (s += c), 0);
      const rand = Math.floor(Math.random() * sumCnt + 1);

      let sumCum = 0;
      for (const [word, count] of Object.entries(node)) {
        sumCum += count;
        if (sumCum > rand) {
          nextWord = word;
          break;
        }
      }

      curLimit++;
    }
  }
}

c1 = new MarkovChain();
c1.modelChain(`In probability theory and statistics, a Markov chain or Markov process is a stochastic process describing a sequence of possible events in which the probability of each event depends only on the state attained in the previous event. Informally, this may be thought of as, "What happens next depends only on the state of affairs now." A countably infinite sequence, in which the chain moves state at discrete time steps, gives a discrete-time Markov chain (DTMC). A continuous-time process is called a continuous-time Markov chain (CTMC). Markov processes are named in honor of the Russian mathematician Andrey Markov.

Markov chains have many applications as statistical models of real-world processes. They provide the basis for general stochastic simulation methods known as Markov chain Monte Carlo, which are used for simulating sampling from complex probability distributions, and have found application in areas including Bayesian statistics, biology, chemistry, economics, finance, information theory, physics, signal processing, and speech processing.

The adjectives Markovian and Markov are used to describe something that is related to a Markov process.`);
c1.generateText("A", 5);
