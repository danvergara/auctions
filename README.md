# **Code Challenge: Auction**

## Overview

This exercise serves as a demonstration of how you solve problems using *idiomatic* **Golang**.

## Rules

There are a few guidelines to follow while completing this exercise:

- **Keep It Confidential:** Please keep this document, the problem, and your solution confidential between yourself and
  One Advisory (d.b.a. Dispatch), even after leaving the interview
- **48-hour Time Limit:** This should be a relatively short but fun exercise to show of your knowledge about specific
  technologies and software engineering principles. Unless otherwise arranged, make handoff to the Dispatch engineering
  team within 48 hours
- **Use Generally Available Tools:** The team evaluates new releases of tools and technologies as they become available,
  however for this exercise, please only use the latest version of tools and technologies that are considered Generally
  Available.
- **Open Documentation, Open Internet:** Engineers spend a non-trivial amount of time sourcing information online; feel
  free to use online resources and be able to demonstrate why you selected one approach to solving a problem over
  another
- **Deliver a Complete Project:** The solution is delivered via a pull request on the repository. It is recommended to
  clong the gist repository to your own Github repo and create a branch for your implementation on your branch. Share
  the pull request to main with us when completed.
- We want to be able to run your solution on our machines

All software projects require careful balance of both technical requirements along with functional requirements. Your
solution must comply with the following technical objectives:

- Implement your solution using a current released version of Go (ie 1.22 or 1.21)
- Do not use 3rd party libraries. Use only the core, standard library
- Must include unit tests to verify your components and include these tests in your solution
- Do not include a graphical user interface of any kind. The quality of code is evaluated only
- We are looking for a **clean**, **well-factored**, idiomatic codebase that has accompanying unit tests
- We want to see how you reason and approach a problem with ambiguity

# Problem Statement / Feature Description

Consider a new and different computerized auction site where a seller can offer an item up for sale and people can bid
against each other to buy the item. The company building this site has asked you to come up with the algorithm to
automatically determine the winning bid after all bidders have entered their information on the site. Your library will
be integrated into the site by other developers working on the project.

## Requirements

- **Starting bid** The first and lowest bid the buyer is willing to offer for the item
- **Max bid** This maximum amount the bidder is willing to pay for the item.
- **Auto-increment amount** A dollar amount that the computer algorithm will add to the bidder's current bid each time
  the bidder is in a losing position relative to the other bidders. The algorithm should never let the current bid
  exceed the Max bid. The algorithm should only allow increments of the exact auto-increment amount
- The amount of the winner's bid should be the lowest amount possible (given all the previous rules) that will win the
  auction.
- Ties are resolved by prioritizing earlier entries.

## Examples

Answers are not provided intensionally, but please include these in your unit tests with what should be the correct
results

### Auction #1

|               | Sasha  | John   | Pat    |
|---------------|--------|--------|--------|
| Initial Bid   | $50.00 | $60.00 | $55.00 |
| Max Bid       | $80.00 | $82.00 | $85.00 |
| Bid Increment | $3.00  | $2.00  | $5.00  |

### Auction #2

|               | Riley   | Morgan  | Charlie |
|---------------|---------|---------|---------|
| Initial Bid   | $700.00 | $599.00 | $625.00 |
| Max Bid       | $725.00 | $725.00 | $725.00 |
| Bid Increment | $2.00   | $15.00  | $8.00   |

### Auction #3

|               | Alex     | Jesse    | Drew     |
|---------------|----------|----------|----------|
| Initial Bid   | $2500.00 | $2800.00 | $2501.00 |
| Max Bid       | $3000.00 | $3100.00 | $3200.00 |
| Bid Increment | $500.00  | $201.00  | $247.00  |

Good Luck!
