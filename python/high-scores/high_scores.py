def latest(scores):
    return scores[-1]


def personal_best(scores):
    return max(scores)


def personal_top_three(scores):
    #* Using sorted or sort they have a reverse flag
    #* These are the two approaches for the solution
    #scores = sorted(scores, reverse=True)
    scores.sort(reverse=True)
    return scores[:3]
